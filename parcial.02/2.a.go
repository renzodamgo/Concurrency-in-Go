package main

import (
    "fmt"
    "sync"
)

var turno = "hombres"
var hombres = 0
var mujeres = 0
var auxiliar = make([]bool, 10)

func hombre(id int, dni string, m *sync.Mutex) {
    for {
        m.Lock()

        if turno == "hombres" && !auxiliar[id] {
            hombres++
            auxiliar[id] = true
            fmt.Printf("Hombre con DNI:%s esta entrado al Rotor.\n", dni)
        }

        if hombres > 9 {
            fmt.Println()
            hombres = 0
            auxiliar = make([]bool, 10)
            turno = "mujeres"
        }

        m.Unlock()
    }
}

func mujer(id int, dni string, m *sync.Mutex) {
    for {
        m.Lock()

        if turno == "mujeres" && !auxiliar[id] {
            mujeres++
            auxiliar[id] = true
            fmt.Printf("Mujer con DNI:%s esta entrado al Rotor.\n", dni)
        }

        if mujeres > 9 {
            fmt.Println()
            mujeres = 0
            auxiliar = make([]bool, 10)
            turno = "hombres"
        }

        m.Unlock()
    }
}

func main() {
    m := new(sync.Mutex)

    hombresDni := [10]string{"42608666", "42389192", "59355072", "91613288", "70986129", "10503611", "66406087", "26057179", "60084725", "73467559"}
    mujeresDni := [10]string{"04802391", "61134033", "89123994", "12825133", "18986160", "41067735", "63580332", "42728669", "91859108", "93338831"}

    for id := 0; id < 10; id++ {
        go hombre(id, hombresDni[id], m)
        go mujer(id, mujeresDni[id], m)
    }

    var esperar string
    fmt.Scanln(&esperar)
}