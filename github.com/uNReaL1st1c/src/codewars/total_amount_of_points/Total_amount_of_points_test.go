package totalamountofpoints

import "testing"

func TestPoints(t *testing.T) {

	firstTeamResult := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	points_first := Points(firstTeamResult)

	if points_first != 30 {
		t.Errorf("incorrect result: expected %v, got %v", 30, points_first)
	}

	secondTeamResult := []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}
	points_second := Points(secondTeamResult)

	if points_second != 10 {
		t.Errorf("incorrect result: expected %v, got %v", 10, points_second)
	}

	thirdTeamResult := []string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}
	points_third := Points(thirdTeamResult)

	if points_third != 0 {
		t.Errorf("incorrect result: expected %v, got %v", 0, points_third)
	}

	fourthTeamResult := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}
	points_fourth := Points(fourthTeamResult)

	if points_fourth != 15 {
		t.Errorf("incorrect result: expected %v, got %v", 15, points_fourth)
	}

	fifthTeamResult := []string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"}
	points_fifth := Points(fifthTeamResult)

	if points_fifth != 12 {
		t.Errorf("incorrect result: expected %v, got %v", 12, points_fifth)
	}

}

func Test_countOfPoints(t *testing.T) {

	firstTeamResult := []string{"1", "0"}
	game_point_first := countOfPoints(firstTeamResult)

	if game_point_first != 3 {
		t.Errorf("incorrect result: expected %v, got %v", 3, game_point_first)
	}

	secondTeamResult := []string{"0", "1"}
	game_point_second := countOfPoints(secondTeamResult)

	if game_point_second != 0 {
		t.Errorf("incorrect result: expected %v, got %v", 0, game_point_second)
	}

	thirdTeamResult := []string{"0", "0"}
	game_point_third := countOfPoints(thirdTeamResult)

	if game_point_third != 1 {
		t.Errorf("incorrect result: expected %v, got %v", 1, game_point_third)
	}

}
