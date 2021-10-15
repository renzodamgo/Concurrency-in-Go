	switch (t->back) {
	default: Uerror("bad return move");
	case  0: goto R999; /* nothing to undo */

		 /* PROC Turnomujeres */
;
		;
		;
		;
		
	case 5: // STATE 3
		;
		now.mujeres = trpt->bup.oval;
		;
		goto R999;
;
		;
		;
		;
		
	case 8: // STATE 6
		;
		now.genero = trpt->bup.oval;
		;
		goto R999;

	case 9: // STATE 7
		;
		now.mujeres = trpt->bup.oval;
		;
		goto R999;

	case 10: // STATE 13
		;
		p_restor(II);
		;
		;
		goto R999;

		 /* PROC Turnohombres */
;
		;
		;
		;
		
	case 13: // STATE 3
		;
		now.hombres = trpt->bup.oval;
		;
		goto R999;
;
		;
		;
		;
		
	case 16: // STATE 6
		;
		now.genero = trpt->bup.oval;
		;
		goto R999;

	case 17: // STATE 7
		;
		now.hombres = trpt->bup.oval;
		;
		goto R999;

	case 18: // STATE 13
		;
		p_restor(II);
		;
		;
		goto R999;
	}

