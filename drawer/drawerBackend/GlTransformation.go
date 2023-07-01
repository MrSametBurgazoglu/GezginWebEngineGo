package drawerBackend

func getRealValue(length, value int) float32 { //transform device position to opengl value
	x := 2.0 * float32(value) / float32(length)
	return -1.0 + x
}
