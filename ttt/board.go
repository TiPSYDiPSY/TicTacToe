package ttt

type board [3][3] string

func (b *board) emptyCount() int {
	count := 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] == "_" {
				count++
			}
		}
	}
	return count
}

func (b *board) winner() string {
	xStreak := [3]string{"X","X","X"}
	oStreak := [3]string{"O","O","O"}

	rows := b.rows()
	for i := range rows {
		if rows[i] == xStreak {
			return "X"
		}
		if rows[i] == oStreak {
			return "O"
		}
	}

	columns := b.coluns()
	for i := range columns {
		if columns[i] == xStreak {
			return "X"
		}
		if columns[i] == oStreak {
			return "O"
		}
	}

	diagonals := b.diagonals()
	for i := range diagonals {
		if diagonals[i] == xStreak {
			return "X"
		}
		if diagonals[i] == oStreak {
			return "O"
		}
	}

}

func (b *board) rows() [3][3]string {
	return *b
}

func (b *board) coluns() [3][3]string {
	colums := [3][3]string{}
	for i := range b {
		for j := range b[i] {
			colums[j][i] = b[i][j]
		}
	}
	return colums
}


func (b *board) diagonals() [2][3]string {
	diagonals := [2][3]string{}
	for i:= range b {
		diagonals[0][i] = b[i][i]
	}

	for i:= range b {
		diagonals[1][i] = b[i][2-i]
	}

	return diagonals
}