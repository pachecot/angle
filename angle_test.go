package angle

import (
	"math"
	"reflect"
	"testing"
)

const epsilon = 0.000001

func equalFloat(got, want float64) bool {
	if got == want {
		return true
	}
	if want == 0 {
		return math.Abs(got) < epsilon
	}
	return math.Abs(got-want) < math.Abs(want)*epsilon
}

func TestDegrees(t *testing.T) {
	type args struct {
		degrees float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		{"Degrees(45)", args{45}, Angle(math.Pi / 4)},
		{"Degrees(90)", args{90}, Angle(math.Pi / 2)},
		{"Degrees(180)", args{180}, Angle(math.Pi)},
		{"Degrees(360)", args{360}, Angle(math.Pi * 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Degrees(tt.args.degrees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Degrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Degrees(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Degrees() = 90", Angle(math.Pi / 2), 90},
		{"Degrees() = 180", Angle(math.Pi), 180},
		{"Degrees() = 360", Angle(math.Pi * 2), 360},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Degrees(); got != tt.want {
				t.Errorf("Angle.Degrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadians(t *testing.T) {
	type args struct {
		radians float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		{"Radians(1)", args{1}, Angle(1)},
		{"Radians(Pi)", args{math.Pi}, Degrees(180)},
		{"Radians(2*Pi)", args{math.Pi * 2}, Degrees(360)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Radians(tt.args.radians); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Radians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinutes(t *testing.T) {
	type args struct {
		minutes float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		{"Minutes(4)", args{4}, Degrees(1.0)},
		{"Minutes(1)", args{1}, Degrees(1.0 / 4.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Minutes(tt.args.minutes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Minutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Radians(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Radians() = 1.0", Radians(1.0), 1.0},
		{"Radians() = Pi", Radians(math.Pi), math.Pi},
		{"Radians() = Pi/2", Degrees(90), math.Pi / 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Radians(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Radians() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Arcminute(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Minutes() = 60", Degrees(1), 60},
		{"Minutes() = 1", Degrees(1.0 / 60.0), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Arcminute(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Arcminute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Minute(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Minute() = 4", Degrees(1.0), 4},
		{"Minute() = 1", Degrees(1.0 / 4.0), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Minutes(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Minute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Hour(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Hour() = 2", Degrees(30.0), 2.0},
		{"Hour() = 1", Degrees(15.0), 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Hours(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Hour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Day(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Day() = 0.5", Degrees(180.0), 0.5},
		{"Day() = 1", Degrees(360.0), 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Days(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Sin(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Sin() = 0.0", Degrees(0), 0.0},
		{"Sin() = 1.0", Degrees(90), 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Sin(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Sin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Cos(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Cos() = 1.0", Degrees(0), 1.0},
		{"Cos() = 0.0", Degrees(90), 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Cos(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Cos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_Tan(t *testing.T) {
	tests := []struct {
		name string
		ang  Angle
		want float64
	}{
		{"Tan() = 0.0", Degrees(0), 0.0},
		{"Tan() = 1.0", Degrees(45), 1.0},
		{"Tan() = -1.0", Degrees(-45), -1.0},
		{"Tan() = +inf", Degrees(90), 1.6331239353195392e+16},
		{"Tan() = +inf", Angle(math.Pi / 2), 1.6331239353195392e+16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ang.Tan(); !equalFloat(got, tt.want) {
				t.Errorf("Angle.Tan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsin(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Asin(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Asin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcos(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Acos(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Acos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtan(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atan(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Atan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtan2(t *testing.T) {
	type args struct {
		y float64
		x float64
	}
	tests := []struct {
		name string
		args args
		want Angle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atan2(tt.args.y, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Atan2() = %v, want %v", got, tt.want)
			}
		})
	}
}
