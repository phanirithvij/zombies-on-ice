package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// Controllers state of the controllers
type Controllers struct {
	All       []ControllerUpdater
	Unplugged []ControllerUpdater
	Plugged   []*PluggedController

	Removed map[*Controller]bool
}

// PluggedController a pluggedin controller
type PluggedController struct {
	Controller *Controller
	Updater    ControllerUpdater
}

// NewControllers a new instance of controllers
func NewControllers() *Controllers {
	controllers := &Controllers{
		All: []ControllerUpdater{
			&Keyboard0,
			&Keyboard1,

			&Gamepad0,
			&Gamepad1,
			&Gamepad2,
			&Gamepad3,
		},
	}

	for _, updater := range controllers.All {
		controllers.Unplugged = append(controllers.Unplugged, updater)
	}

	return controllers
}

// Update updates the controllers
func (controllers *Controllers) Update(window *glfw.Window) {
	plugged := []*PluggedController{}
	unplugged := []ControllerUpdater{}
	removed := map[*Controller]bool{}

	for _, plug := range controllers.Plugged {
		plug.Updater.Update(plug.Controller, window)
		if plug.Controller.Connected && !plug.Controller.Back {
			plugged = append(plugged, plug)
		} else {
			unplugged = append(unplugged, plug.Updater)
			removed[plug.Controller] = true
			log.Println("Removed controller")
		}
	}

	for _, unplug := range controllers.Unplugged {
		temp := &Controller{}
		unplug.Update(temp, window)
		if temp.Connected && temp.Active() {
			plugged = append(plugged, &PluggedController{nil, unplug})
			log.Println("Added controller")
		} else {
			unplugged = append(unplugged, unplug)
		}
	}

	controllers.Plugged = plugged
	controllers.Unplugged = unplugged
	controllers.Removed = removed
}

var (
	// Keyboard0 the main arrow keyboard
	Keyboard0 = Keyboard{
		Up:    glfw.KeyUp,
		Down:  glfw.KeyDown,
		Left:  glfw.KeyLeft,
		Right: glfw.KeyRight,

		Start: glfw.KeyEnter,
		Back:  glfw.KeyBackspace,

		A: glfw.KeyO,
		B: glfw.KeyP,
	}
	// Keyboard1 the main awsd keyboard
	Keyboard1 = Keyboard{
		Up:    glfw.KeyW,
		Down:  glfw.KeyS,
		Left:  glfw.KeyA,
		Right: glfw.KeyD,

		Start: glfw.KeyT,
		Back:  glfw.KeyR,

		A: glfw.KeyTab,
		B: glfw.KeyQ,
	}

	// Gamepad0 the 0th gamepad
	Gamepad0 = Gamepad{glfw.Joystick1, 0.05}
	// Gamepad1 the 1st gamepad
	Gamepad1 = Gamepad{glfw.Joystick2, 0.05}
	// Gamepad2 the 2nd gamepad
	Gamepad2 = Gamepad{glfw.Joystick3, 0.05}
	// Gamepad3 the 3rd gamepad
	Gamepad3 = Gamepad{glfw.Joystick4, 0.05}
)
