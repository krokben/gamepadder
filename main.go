package main

import (
  "github.com/go-vgo/robotgo"
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/platforms/joystick"
)

type Mods struct {
  l2 bool
  r2 bool
}

func (m *Mods) OutputKeyboard(r2 string, l2 string, noMod string) {
  if m.r2 {
    robotgo.KeyTap(r2)
    return
  }
  if m.l2 {
    robotgo.TypeStr(l2)
    return
  }
  robotgo.TypeStr(noMod)
}

func main() {
  joystickAdaptor := joystick.NewAdaptor()
  stick := joystick.NewDriver(joystickAdaptor, "dualshock4")

  mods := &Mods{false, false}

  work := func() {
    // d-pad
    stick.On(joystick.DownRelease, func(data interface{}) {
      mods.OutputKeyboard("down", "b", "n")
    })
    stick.On(joystick.LeftRelease, func(data interface{}) {
      mods.OutputKeyboard("left", "l", "o")
    })
    stick.On(joystick.UpRelease, func(data interface{}) {
      mods.OutputKeyboard("up", "u", "s")
    })
    stick.On(joystick.RightRelease, func(data interface{}) {
      mods.OutputKeyboard("right", "c", "h")
    })

    // buttons
    stick.On(joystick.XRelease, func(data interface{}) {
      mods.OutputKeyboard("", "m", "e")
    })
    stick.On(joystick.SquareRelease, func(data interface{}) {
      mods.OutputKeyboard("", "f", "t")
    })
    stick.On(joystick.TriangleRelease, func(data interface{}) {
      mods.OutputKeyboard("", "w", "a")
    })
    stick.On(joystick.CircleRelease, func(data interface{}) {
      mods.OutputKeyboard("", "y", "i")
    })
    stick.On(joystick.L1Release, func(data interface{}) {
      mods.OutputKeyboard("", "g", "r")
    })
    stick.On(joystick.R1Release, func(data interface{}) {
      mods.OutputKeyboard("", "p", "d")
    })

    // utility
    stick.On(joystick.ShareRelease, func(data interface{}) {
      mods.OutputKeyboard("", "j", "q")
    })
    stick.On(joystick.OptionsRelease, func(data interface{}) {
      mods.OutputKeyboard("", "x", "k")
    })
    stick.On(joystick.HomeRelease, func(data interface{}) {
      mods.OutputKeyboard("", "z", "v")
    })
    stick.On(joystick.L3Release, func(data interface{}) {
      robotgo.KeyTap("backspace")
    })
    stick.On(joystick.R3Release, func(data interface{}) {
      if mods.r2 {
        robotgo.KeyTap("enter")
        return
      }
      robotgo.KeyTap("space")
    })

    // modifiers
    stick.On(joystick.L2Press, func(data interface{}) {
      mods.l2 = true
    })
    stick.On(joystick.L2Release, func(data interface{}) {
      mods.l2 = false
    })
    stick.On(joystick.R2Press, func(data interface{}) {
      mods.r2 = true
    })
    stick.On(joystick.R2Release, func(data interface{}) {
      mods.r2 = false
    })
  }

  robot := gobot.NewRobot("joystickBot",
    []gobot.Connection{joystickAdaptor},
    []gobot.Device{stick},
    work,
  )

  robot.Start()
}
