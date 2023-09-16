package drawerBackend

type Rect interface {
	GetPosition() (int, int)
	GetSize() (int, int)
}
