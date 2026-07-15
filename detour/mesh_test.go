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

func TestValidateLinksRejectsInvalidPolygonReference(t *testing.T) {
	var mesh NavMesh
	if status := mesh.Init(&NavMeshParams{MaxTiles: 1, MaxPolys: 8}); StatusFailed(status) {
		t.Fatalf("Init failed: status=0x%x", status)
	}
	mesh.Tiles[0].Header = &MeshHeader{PolyCount: 1}
	mesh.Tiles[0].Polys = []Poly{{FirstLink: 0}}
	mesh.Tiles[0].Links = []Link{{Ref: mesh.encodePolyID(1, 0, 1), Next: nullLink}}
	if err := mesh.ValidateLinks(); err == nil {
		t.Fatal("ValidateLinks succeeded with an out-of-range polygon reference")
	}
}
