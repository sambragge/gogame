package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type World struct {
	undead []*Undead
}

func (w *World) collides(entity Undead)bool{


	for _, other := range w.undead {
		if !other.IsEqual(entity) {
			if CheckCollisionRecs(other.Rectangle, entity.Rectangle) {
				return true
			}
		}
	}

	return false
}

func (w *World) playerOnField()bool{

	for _, u := range w.undead {
		if u.Possessed() {
			return true
		}
	}

	return false
}

func (w *World) gameStillGoing()bool{

	//leftAlive := 0

	//for _, u := range w.undead {
	//	if u.alive() {
	//		leftAlive++
	//	}
	//}

	if !w.playerOnField() {
		return false
	}
	return true
}

func (w *World) player() *Undead {
	for _, u := range w.undead {
		if u.Possessed() {
			return u
		}
	}
	return nil
}

func newWorld() *World {

	enemyCount := 5
	w := &World{}


	w.addUndead(NewUndead("player", w))
	for i := 0; i < enemyCount; i++ {
		w.addUndead(NewUndead("hostile", w))
	}




	return w
}

func (w *World) _update(){
	if w.gameStillGoing() {
		for _, u := range w.undead {
			u.Update()
		}
	}else{
		DrawText("GAME OVER", 100, 250, 40, Black)
	}
}

func (w *World) _draw(){
	for _, u := range w.undead {
		u.Draw()
	}
}

func (w *World) entityHit(pos Vector2)(bool, *Undead){

	for _, u := range w.undead {
		if u.isHit(pos) {

			return true, u
		}
	}

	return false, nil
}

func(w *World) draw(){
	BeginDrawing()

	ClearBackground(RayWhite)

	w._update()
	w._draw()

	EndDrawing()
}

func(w *World) addUndead(u *Undead){
	w.undead = append(w.undead, u)
}