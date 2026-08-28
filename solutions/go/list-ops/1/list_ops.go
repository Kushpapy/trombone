package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
    for _, v := range s {
        acc = fn(acc, v)
    }

    return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
    for i := len(s) - 1; i >= 0; i-- {
        acc = fn(s[i], acc)
    }

    return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	 

    count := 0
    for _,v := range s {
        if fn(v){
            count++
        }
    }

    initLength := 0
    newList := make([]int, count)
    for _, v := range s {
        if fn(v) {
            newList[initLength] = v
            initLength += 1
        }
        
    }

    return newList
}

func (s IntList) Length() int {
	count := 0

    for range s {
        count += 1
    }

    return count
}

func (s IntList) Map(fn func(int) int) IntList {
	out := make(IntList,len(s))
    for i, v := range s {
        out[i] = fn(v)
    }
    return out
}

func (s IntList) Reverse() IntList {
	begin := 0
    end := s.Length() - 1

    newArr := make([]int, s.Length())
    
    for begin <= end {
       newArr[end] = s[begin] 
	   newArr[begin] = s[end]
        
        begin++
        end--
    }

    return newArr
}

func (s IntList) Append(lst IntList) IntList {
    total := len(s) + len(lst) 
	newSlice := make([]int, total)

    count := 0
    for i,v := range s {
        newSlice[i] = v
        count++
    }

    for _, v := range lst {
        newSlice[count] = v
        count++
    }

    return newSlice
}

func (s IntList) Concat(lists []IntList) IntList {
	total := len(s)

    for _, l := range lists {
        total += len(l)
    }

    out := make(IntList, total)
    index := 0
    for _, v := range s {
        out[index] = v
        index++
    }
    for _, l := range lists {
        for _, v := range l {
            out[index] = v
            index++
        }
    }

    return out
}
