package main

import (
	"log"
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	log.Println("Start main tests")
	res := m.Run()
	log.Println("Stop main tests")
	
	os.Exit(res)
	
}


func TestAdd(t *testing.T) {
	t.Log("TestAdd")

	t.Run("simple", func(t *testing.T) {
		t.Log("TestAdd-simple")

		var x, y, res int = 2, 3, 5
		real_res := Add(x, y)

		if res != real_res {
			t.Errorf("(res)%d != (real_res)%d", res, real_res)
		}

	})

	t.Run("medium", func(t *testing.T) {
		t.Log("TestAdd-medium")

		var x, y, res int = 1982, -951, 1031
		real_res := Add(x, y)

		if res != real_res {
			t.Errorf("(res)%d != (real_res)%d", res, real_res)
		}
		
	})

	t.Run("hard", func(t *testing.T) {
		t.Log("TestAdd-hard")

		var x, y, res int = -4_532_453_535, -3_546_345, -4_535_999_880
		real_res := Add(x, y)

		if res != real_res {
			t.Errorf("(res)%d != (real_res)%d", res, real_res)
		}
		
	})

}


func TestMulti(t *testing.T) {
	t.Log("TestMulti")

	t.Run("simple", func(t *testing.T) {
		t.Log("TestMulti-simple")

		var x, y, res int = 2, 3, 6
		real_res := Multi(x, y)

		if res != real_res {
			t.Errorf("(res)%d != (real_res)%d", res, real_res)
		}

	})

	t.Run("medium", func(t *testing.T) {
		t.Log("TestMulti-medium")

		var x, y, res int = 1982, -951, -1_884_882
		real_res := Multi(x, y)

		if res != real_res {
			t.Errorf("(res)%d != (real_res)%d", res, real_res)
		}
		
	})

	t.Run("hard", func(t *testing.T) {
		t.Log("TestMulti-hard")

		var x, y, res int = -4535, -6345, 28_774_575
		real_res := Multi(x, y)

		if res != real_res {
			t.Errorf("(res)%d != (real_res)%d", res, real_res)
		}
		
	})
	
}
