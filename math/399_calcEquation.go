package main

import "fmt"

//// map记录单个除法结果，但是没法记录带传递的
//func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
//	var res []float64
//	length := len(equations)
//	m := make(map[string]float64, length)
//	for i := 0; i < length; i++ {
//		key := equations[i][0] + "_" + equations[i][1]
//		m[key] = values[i]
//	}
//	for i := 0; i < length; i++ {
//		key := queries[i][0] + "_" + queries[i][1]
//		if _, ok := m[key]; ok {
//			res = append(res, m[key])
//		} else {
//			_key := queries[i][1] + "_" + queries[i][0]
//			if _, ok = m[_key]; ok {
//				res = append(res, 1/m[_key])
//			} else {
//				res = append(res, -1.0)
//			}
//		}
//	}
//	return res
//}

//// 模拟除法，给出所有变量的值，缺点：没有解决变量之间的关联问题
//func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
//	length := len(equations)
//	var res []float64
//	m := make(map[string]float64)
//	for i := 0; i < length; i++ {
//		_, ok1 := m[equations[i][0]]
//		_, ok2 := m[equations[i][1]]
//		if !ok1 && !ok2 {
//			m[equations[i][0]] = 1.0
//			m[equations[i][1]] = 1.0 / values[i]
//		} else if !ok1 {
//			m[equations[i][0]] = m[equations[i][1]] * values[i]
//		} else if !ok2 {
//			m[equations[i][1]] = m[equations[i][0]] / values[i]
//		}
//	}
//	for i := 0; i < len(queries); i++ {
//		val1, ok1 := m[queries[i][0]]
//		val2, ok2 := m[queries[i][1]]
//		if ok1 && ok2 {
//			res = append(res, val1/val2)
//		} else {
//			res = append(res, -1.0)
//		}
//	}
//	return res
//}

// 模拟除法，尝试改进，但是没法用赋值的方法解决变量关联的问题
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var res []float64
	m := make(map[string]float64)

	for len(equations) > 0 {
		// 初始化，给第一组元素赋值
		_, ok1 := m[equations[0][0]]
		_, ok2 := m[equations[0][1]]
		if !ok1 && !ok2 {
			m[equations[0][0]] = 1.0
			m[equations[0][1]] = 1.0 / values[0]
			equations = equations[1:]
			values = values[1:]
		}
		// 遍历，解出所有与第一组相关联的变量，剩余的留在equations中
		for i := 0; i < len(equations); i++ {
			_, ok1 := m[equations[i][0]]
			_, ok2 := m[equations[i][1]]
			if !ok1 && !ok2 {
				continue
			} else if !ok1 {
				m[equations[i][0]] = m[equations[i][1]] * values[i]
				equations = append(equations[:i], equations[i+1:]...)
				values = append(values[:i], values[i+1:]...)
			} else if !ok2 {
				m[equations[i][1]] = m[equations[i][0]] / values[i]
				equations = append(equations[:i], equations[i+1:]...)
				values = append(values[:i], values[i+1:]...)
			}
		}
	}
	for i := 0; i < len(queries); i++ {
		val1, ok1 := m[queries[i][0]]
		val2, ok2 := m[queries[i][1]]
		if ok1 && ok2 {
			res = append(res, val1/val2)
		} else {
			res = append(res, -1.0)
		}
	}
	return res
}

func main() {
	d1 := [][]string{
		[]string{"a", "b"},
		[]string{"e", "f"},
		[]string{"b", "e"},
	}
	values := []float64{3.4, 1.4, 2.3}
	d2 := [][]string{
		[]string{"b", "a"},
		[]string{"a", "f"},
	}
	fmt.Println(calcEquation(d1, values, d2))
}
