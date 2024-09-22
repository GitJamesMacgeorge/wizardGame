package main 

type Entity interface {
  MoveUp()
  MoveDown()
  MoveRight()
  MoveLeft()
}

func (p *Player) MoveUp() {
  p.y -= 2
}

func MoveUp(e Entity) {
  e.MoveUp()
}

func (p *Player) MoveDown() {
  p.y += 2
}

func MoveDown(e Entity) {
  e.MoveDown()
}

func (p *Player) MoveRight() {
  p.x += 2
}

func MoveRight(e Entity) {
  e.MoveRight()
}

func (p *Player) MoveLeft() {
  p.x -= 2
}

func MoveLeft(e Entity) {
  e.MoveLeft()
}


