byte hombres = 0
byte mujeres = 0
byte genero = 0

active proctype Turnohombres() {
    do
    ::
        (genero == 0)
            if
            :: (hombres <= 9) -> hombres++
            printf("Hombre %d entrando a rotor \n ", hombres)
            :: (hombres > 9) -> genero = 1
            hombres = 0
            fi 
    od
}

active proctype Turnomujeres() {
    do
    ::
        (genero == 1)
            if
            :: (mujeres <= 9) -> mujeres++
             printf("Mujer %d entrando a rotor \n ", mujeres)
            :: (mujeres > 9) -> genero = 0
            mujeres = 0
            fi 
    od
}