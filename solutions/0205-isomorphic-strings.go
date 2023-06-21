// Day 11

func isIsomorphic(s string, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}
    for i := 0; i < len(t); i++ {
        v1 := s[i]
        v2 := t[i]
        if vv1, ok := m1[v1]; ok {
            if vv1 != t[i] {
                return false
            }
        } else {
            m1[v1] = v2
        }
        if vv2, ok := m2[v2]; ok {
            if vv2 != s[i] {
                return false
            }
        } else {
            m2[v2] = v1
        }
    }
    return true
}