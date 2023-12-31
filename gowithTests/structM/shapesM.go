package shape333

import("math")
//---------------------
type Rectangle struct {
  Width  float64
  Height float64
}
//----
// define method func Area()
func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}
//-----------------------------------------

type Circle struct {
  Radius float64
}
//-----
// define method func Area()
func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

