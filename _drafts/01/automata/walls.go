package automata

type WallPosition byte

const (
	WallPositionTop WallPosition = iota
	WallPositionBottom
	WallPositionLeft
	WallPositionRight
)

func (wp WallPosition) GetOpposite() WallPosition {
	switch wp {
	case WallPositionTop:
		return WallPositionBottom
	case WallPositionBottom:
		return WallPositionTop
	case WallPositionLeft:
		return WallPositionRight
	default:
		return WallPositionLeft
	}
}

type WallEffect byte

const (
	WallEffectNone WallEffect = iota
	WallEffectReverseDirection
	WallEffectTeleportOpposite
)

type Wall struct {
	position WallPosition
	reaction WallEffect
}

func NewWall(position WallPosition, reaction WallEffect) *Wall {
	return &Wall{
		position: position,
		reaction: reaction,
	}
}
