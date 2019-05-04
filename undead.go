package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	"gopkg.in/mgo.v2/bson"
)

type Undead struct {
	Rectangle
	*World
	ID bson.ObjectId
	color     Color
	speed     float32
	possessed bool
	hostile bool
	deadDead bool
	health float32
}

func NewUndead(kind string, world *World) *Undead {

	switch(kind){

	case "player":
		u := &Undead{
			NewRectangle(100, 100, 50, 50),
			world,
			bson.NewObjectId(),
			Blue,
			5.0,
			true,
			false,
			false,
			10.0,
		}
		return u


	default:
		u := &Undead{
			NewRectangle(float32(GetRandomValue(0, 800)), float32(GetRandomValue(0, 500)), 50, 50),
			world,
			bson.NewObjectId(),
			Green,
			1.0,
			false,
			true,
			false,
			10.0,
		}
		return u

	}

}

func (u *Undead) takeDamage(dmg float32){
	u.health -= dmg
	healthString := fmt.Sprintf("%f", u.health)
	DrawText(healthString,  int32(u.X), int32(u.Y-u.Height), 30, Black)
}

func (u *Undead) ChangeColor(color Color){
	u.color = color
}


func (u *Undead) posAsVec2()Vector2{
	return NewVector2(u.X, u.Y)
}

func(u *Undead) Possessed()bool{
	return u.possessed
}

func (u *Undead) IsEqual(other Undead)bool{
	return u.ID == other.ID
}

func(u *Undead) _playerUpdate(){
	if u.color != Blue { u.color = Blue }
	if IsKeyDown(KeyA) {
		u.move("left")
	}
	if IsKeyDown(KeyW) {
		u.move("up")
	}
	if IsKeyDown(KeyS) {
		u.move("down")
	}
	if IsKeyDown(KeyD) {
		u.move("right")
	}

	// mouse inputs
	if IsMouseButtonDown(MouseLeftButton) {
		u.tryToPossess(GetMousePosition())
	}
}

func(u *Undead) _hostileUpdate(){
	if u.color != Green { u.color = Green }
	if u.player() != nil {
		u.moveTowardsPlayer()
	}
}

func (u *Undead) moveTowardsPlayer(){

	p := u.player()

	if u.Y < p.Y {
		u.move("down")
	}else if u.Y > p.Y {
		u.move("up")
	}
	if u.X < p.X {
		u.move("right")
	}else if u.X > p.X {
		u.move("left")
	}
}


func(u *Undead) getPhysicsBody()[]Vector2{
	slice := make([]Vector2, 0)

	for y := u.Y; y <= u.Y+u.Height; y++ {
		for x := u.X; x <= u.X+u.Width; x++ {
			slice = append(slice, NewVector2(x, y))
		}
	}

	return slice
}

func(u *Undead) isHit(pos Vector2) bool {


	for _, vec := range u.getPhysicsBody() {
		if pos.X == vec.X && pos.Y == vec.Y {
			u.ChangeColor(Red)
			return true
		}

	}

	return false
}

func (u *Undead) isHostile()bool{
	return u.hostile
}


func(u *Undead) possess(other *Undead){
	u.possessed, other.possessed = false, true
	u.hostile, other.hostile = true, false
	u.speed, other.speed = 1.0, 5.0
}

func(u *Undead) tryToPossess(pos Vector2){


	hit, entity := u.entityHit(pos)
	if hit {
		entity.ChangeColor(Red)
		if !entity.Possessed() {
			u.possess(entity)
		}
	}

}

func (u *Undead) move(direction string){
	switch(direction){
	case "up":
		if !u.collides(u.ifMoved("up")){
			u.Y -= u.speed
		}else{
			u.takeDamage(0.1)
		}
		break

	case "down":
		if !u.collides(u.ifMoved("down")){
			u.Y += u.speed
		}else{
			u.takeDamage(0.1)
		}
		break

	case "left":
		if !u.collides(u.ifMoved("left")) {
			u.X -= u.speed
		}else{
			u.takeDamage(0.1)
		}
		break

	case "right":
		if !u.collides(u.ifMoved("right")){
			u.X += u.speed
		}else{
			u.takeDamage(0.1)
		}
		break
	}
}

func(u *Undead) ifMoved(direction string)Undead{
	temp := *u

	switch(direction){
	case "up":
		temp.Rectangle.Y -= temp.speed
		break

	case "down":
		temp.Rectangle.Y += temp.speed
		break

	case "left":
		temp.Rectangle.X -= temp.speed
		break

	case "right":
		temp.Rectangle.X += temp.speed
		break
	}

	return temp
}

func(u *Undead) alive()bool{
	return !u.deadDead
}

func (u *Undead) die(){
	u.possessed = false
	u.deadDead = true
}


func (u *Undead) Update(){

	if u.health < 0.0 {
		u.die()
	}
	if !u.deadDead {

		if u.possessed {
			u._playerUpdate()
		}else {
			if u.hostile {
				u._hostileUpdate()
			}
		}
	}else{
		u.ChangeColor(Black)
	}
}

func(u *Undead) drawInfo(info string){
	DrawText(info,100, 250, 34, Black)
}

func(u *Undead) Draw(){
	pos := NewVector2(u.X, u.Y)
	size := NewVector2(u.Width, u.Height)
	DrawRectangleV(pos, size, u.color)
}
