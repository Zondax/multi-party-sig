package zklogstar

import (
	"crypto/rand"

	"github.com/cronokirby/safenum"
	"github.com/taurusgroup/multi-party-sig/internal/hash"
	"github.com/taurusgroup/multi-party-sig/pkg/math/arith"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/math/sample"
	"github.com/taurusgroup/multi-party-sig/pkg/paillier"
	"github.com/taurusgroup/multi-party-sig/pkg/pedersen"
)

type Public struct {
	// C = Enc₀(x;ρ)
	// Encryption of x under the prover's key
	C *paillier.Ciphertext

	// X = x⋅G
	X *curve.Point

	// G is the base point of the curve.
	// If G = nil, the default base point is used.
	G *curve.Point

	Prover *paillier.PublicKey
	Aux    *pedersen.Parameters
}

type Private struct {
	// X is the plaintext of C and the discrete log of X.
	X *safenum.Int

	// Rho = ρ is nonce used to encrypt C.
	Rho *safenum.Nat
}

type Commitment struct {
	// S = sˣ tᵘ (mod N)
	S *safenum.Nat
	// A = Enc₀(alpha; r)
	A *paillier.Ciphertext
	// Y = α⋅G
	Y *curve.Point
	// D = sᵃ tᵍ (mod N)
	D *safenum.Nat
}

type Proof struct {
	*Commitment
	// Z1 = α + e x
	Z1 *safenum.Int
	// Z2 = r ρᵉ mod N
	Z2 *safenum.Nat
	// Z3 = γ + e μ
	Z3 *safenum.Int
}

func (p *Proof) IsValid(public Public) bool {
	if p == nil {
		return false
	}
	if !public.Prover.ValidateCiphertexts(p.A) {
		return false
	}
	if p.Y.IsIdentity() {
		return false
	}
	if !arith.IsValidNatModN(public.Prover.N(), p.Z2) {
		return false
	}
	return true
}

func NewProof(hash *hash.Hash, public Public, private Private) *Proof {
	N := public.Prover.N()
	NModulus := public.Prover.Modulus()

	if public.G == nil {
		public.G = curve.NewBasePoint()
	}

	alpha := sample.IntervalLEps(rand.Reader)
	r := sample.UnitModN(rand.Reader, N)
	mu := sample.IntervalLN(rand.Reader)
	gamma := sample.IntervalLEpsN(rand.Reader)

	commitment := &Commitment{
		A: public.Prover.EncWithNonce(alpha, r),
		Y: curve.NewIdentityPoint().ScalarMult(curve.NewScalarInt(alpha), public.G),
		S: public.Aux.Commit(private.X, mu),
		D: public.Aux.Commit(alpha, gamma),
	}

	e, _ := challenge(hash, public, commitment)

	// z1 = α + e x,
	z1 := new(safenum.Int).Mul(e, private.X, -1)
	z1.Add(z1, alpha, -1)
	// z2 = r ρᵉ mod N,
	z2 := NModulus.ExpI(private.Rho, e)
	z2.ModMul(z2, r, N)
	// z3 = γ + e μ,
	z3 := new(safenum.Int).Mul(e, mu, -1)
	z3.Add(z3, gamma, -1)

	return &Proof{
		Commitment: commitment,
		Z1:         z1,
		Z2:         z2,
		Z3:         z3,
	}
}

func (p Proof) Verify(hash *hash.Hash, public Public) bool {
	if !p.IsValid(public) {
		return false
	}

	if public.G == nil {
		public.G = curve.NewBasePoint()
	}

	if !arith.IsInIntervalLPrimeEps(p.Z1) {
		return false
	}

	prover := public.Prover

	e, err := challenge(hash, public, p.Commitment)
	if err != nil {
		return false
	}

	if !public.Aux.Verify(p.Z1, p.Z3, e, p.D, p.S) {
		return false
	}

	{
		// lhs = Enc(z₁;z₂)
		lhs := prover.EncWithNonce(p.Z1, p.Z2)

		// rhs = (e ⊙ C) ⊕ A
		rhs := public.C.Clone().Mul(prover, e).Add(prover, p.A)
		if !lhs.Equal(rhs) {
			return false
		}
	}

	{
		// lhs = [z₁]G
		lhs := curve.NewIdentityPoint().ScalarMult(curve.NewScalarInt(p.Z1), public.G)

		// rhs = Y + [e]X
		eX := curve.NewIdentityPoint().ScalarMult(curve.NewScalarInt(e), public.X)
		rhs := curve.NewIdentityPoint().Add(p.Y, eX)

		if !lhs.Equal(rhs) {
			return false
		}

	}

	return true
}

func challenge(hash *hash.Hash, public Public, commitment *Commitment) (e *safenum.Int, err error) {
	err = hash.WriteAny(public.Aux, public.Prover, public.C, public.X, public.G,
		commitment.S, commitment.A, commitment.Y, commitment.D)
	e = sample.IntervalScalar(hash.Digest())
	return
}
