def isIsomorphic(s: str, t: str) -> bool:
    sToT = {}
    tToS = {}

    for x in range(0, len(s), 1):
        if s[x] in sToT:
            if sToT[s[x]] != t[x]:
                return False
        else:
            sToT[s[x]] = t[x]

        if t[x] in tToS:
            if tToS[t[x]] != s[x]:
                return False
        else:
            tToS[t[x]] = s[x]

    return True


isIsomorphic("paper", "title")
