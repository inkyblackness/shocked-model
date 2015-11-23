package model

type TileProperties struct {
	Type *TileType `json:"type"`

	FloorHeight   *HeightUnit `json:"floorHeight"`
	CeilingHeight *HeightUnit `json:"ceilingHeight"`
	SlopeHeight   *HeightUnit `json:"slopeHeight"`

	RealWorld *RealWorldTileProperties `json:"realWorld"`
	//Cyberspace *CyberspaceTileProperties
}
