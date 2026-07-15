package detour

import "testing"

func TestInitAcceptsSevenSaltBitsForStaticMesh(t *testing.T) {
	var mesh NavMesh
	status := mesh.Init(&NavMeshParams{MaxTiles: 1 << 16, MaxPolys: 1 << 9})
	if StatusFailed(status) {
		t.Fatalf("Init failed with seven salt bits: status=0x%x", status)
	}
}

func TestInitRejectsSixSaltBits(t *testing.T) {
	var mesh NavMesh
	status := mesh.Init(&NavMeshParams{MaxTiles: 1 << 16, MaxPolys: 1 << 10})
	if !StatusFailed(status) {
		t.Fatalf("Init succeeded with six salt bits: status=0x%x", status)
	}
}
