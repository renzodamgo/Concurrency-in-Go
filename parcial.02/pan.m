#define rand	pan_rand
#define pthread_equal(a,b)	((a)==(b))
#if defined(HAS_CODE) && defined(VERBOSE)
	#ifdef BFS_PAR
		bfs_printf("Pr: %d Tr: %d\n", II, t->forw);
	#else
		cpu_printf("Pr: %d Tr: %d\n", II, t->forw);
	#endif
#endif
	switch (t->forw) {
	default: Uerror("bad forward move");
	case 0:	/* if without executable clauses */
		continue;
	case 1: /* generic 'goto' or 'skip' */
		IfNotBlocked
		_m = 3; goto P999;
	case 2: /* generic 'else' */
		IfNotBlocked
		if (trpt->o_pm&1) continue;
		_m = 3; goto P999;

		 /* PROC Turnomujeres */
	case 3: // STATE 1 - 2b.pml:21 - [((genero==1))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][1] = 1;
		if (!((((int)now.genero)==1)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 4: // STATE 2 - 2b.pml:23 - [((mujeres<=9))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][2] = 1;
		if (!((((int)now.mujeres)<=9)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 3 - 2b.pml:23 - [mujeres = (mujeres+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[1][3] = 1;
		(trpt+1)->bup.oval = ((int)now.mujeres);
		now.mujeres = (((int)now.mujeres)+1);
#ifdef VAR_RANGES
		logval("mujeres", ((int)now.mujeres));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 6: // STATE 4 - 2b.pml:24 - [printf('Mujer %d entrando a rotor \\n ',mujeres)] (0:0:0 - 1)
		IfNotBlocked
		reached[1][4] = 1;
		Printf("Mujer %d entrando a rotor \n ", ((int)now.mujeres));
		_m = 3; goto P999; /* 0 */
	case 7: // STATE 5 - 2b.pml:25 - [((mujeres>9))] (0:0:0 - 1)
		IfNotBlocked
		reached[1][5] = 1;
		if (!((((int)now.mujeres)>9)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 8: // STATE 6 - 2b.pml:25 - [genero = 0] (0:0:1 - 1)
		IfNotBlocked
		reached[1][6] = 1;
		(trpt+1)->bup.oval = ((int)now.genero);
		now.genero = 0;
#ifdef VAR_RANGES
		logval("genero", ((int)now.genero));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 9: // STATE 7 - 2b.pml:26 - [mujeres = 0] (0:0:1 - 1)
		IfNotBlocked
		reached[1][7] = 1;
		(trpt+1)->bup.oval = ((int)now.mujeres);
		now.mujeres = 0;
#ifdef VAR_RANGES
		logval("mujeres", ((int)now.mujeres));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 10: // STATE 13 - 2b.pml:29 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[1][13] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */

		 /* PROC Turnohombres */
	case 11: // STATE 1 - 2b.pml:8 - [((genero==0))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		if (!((((int)now.genero)==0)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 12: // STATE 2 - 2b.pml:10 - [((hombres<=9))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][2] = 1;
		if (!((((int)now.hombres)<=9)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 13: // STATE 3 - 2b.pml:10 - [hombres = (hombres+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][3] = 1;
		(trpt+1)->bup.oval = ((int)now.hombres);
		now.hombres = (((int)now.hombres)+1);
#ifdef VAR_RANGES
		logval("hombres", ((int)now.hombres));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 14: // STATE 4 - 2b.pml:11 - [printf('Hombre %d entrando a rotor \\n ',hombres)] (0:0:0 - 1)
		IfNotBlocked
		reached[0][4] = 1;
		Printf("Hombre %d entrando a rotor \n ", ((int)now.hombres));
		_m = 3; goto P999; /* 0 */
	case 15: // STATE 5 - 2b.pml:12 - [((hombres>9))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][5] = 1;
		if (!((((int)now.hombres)>9)))
			continue;
		_m = 3; goto P999; /* 0 */
	case 16: // STATE 6 - 2b.pml:12 - [genero = 1] (0:0:1 - 1)
		IfNotBlocked
		reached[0][6] = 1;
		(trpt+1)->bup.oval = ((int)now.genero);
		now.genero = 1;
#ifdef VAR_RANGES
		logval("genero", ((int)now.genero));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 17: // STATE 7 - 2b.pml:13 - [hombres = 0] (0:0:1 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		(trpt+1)->bup.oval = ((int)now.hombres);
		now.hombres = 0;
#ifdef VAR_RANGES
		logval("hombres", ((int)now.hombres));
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 18: // STATE 13 - 2b.pml:16 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][13] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */
	case  _T5:	/* np_ */
		if (!((!(trpt->o_pm&4) && !(trpt->tau&128))))
			continue;
		/* else fall through */
	case  _T2:	/* true */
		_m = 3; goto P999;
#undef rand
	}

