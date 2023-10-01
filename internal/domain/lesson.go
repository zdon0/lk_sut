package domain

type TimeStamp struct {
	Hour   int
	Minute int
}

type Lesson struct {
	LeftBorder  TimeStamp
	RightBorder TimeStamp
}

func AvailableLessons() []Lesson {
	return []Lesson{
		{
			LeftBorder: TimeStamp{
				Hour:   9,
				Minute: 00,
			},
			RightBorder: TimeStamp{
				Hour:   10,
				Minute: 35,
			},
		},
		{
			LeftBorder: TimeStamp{
				Hour:   10,
				Minute: 45,
			},
			RightBorder: TimeStamp{
				Hour:   12,
				Minute: 20,
			},
		},
		{
			LeftBorder: TimeStamp{
				Hour:   13,
				Minute: 00,
			},
			RightBorder: TimeStamp{
				Hour:   14,
				Minute: 35,
			},
		},
		{
			LeftBorder: TimeStamp{
				Hour:   14,
				Minute: 45,
			},
			RightBorder: TimeStamp{
				Hour:   16,
				Minute: 20,
			},
		},
		{
			LeftBorder: TimeStamp{
				Hour:   16,
				Minute: 30,
			},
			RightBorder: TimeStamp{
				Hour:   18,
				Minute: 05,
			},
		},
	}
}
