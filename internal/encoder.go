package internal

import (
	"strings"
)

type bri struct {
    smap map[string]rune
    rc, rl map[rune]string
}

func NewBri() *bri {
    var rl = map[rune]string{
            '\u2801':"A",
            '\u2803':"B",
            '\u2809':"C",
            '\u2819':"D",
            '\u2811':"E",
            '\u280B':"F",
            '\u281B':"G",
            '\u2813':"H",
            '\u280A':"I",
            '\u281A':"J",
            '\u2805':"K",
            '\u2807':"L",
            '\u280D':"M",
            '\u281D':"N",
            '\u2815':"O",
            '\u280F':"P",
            '\u281F':"Q",
            '\u2817':"R",
            '\u280E':"S",
            '\u281E':"T",
            '\u2825':"U",
            '\u2827':"V",
            '\u283A':"W",
            '\u282D':"X",
            '\u283D':"Y",
            '\u2835':"Z",
            '\u2832':".",
            '\u2816':"!",
            '\u2824':"-",
            '\u2826':"\"",
            '\u2834':"\"",
            '\u2823':"(",
            '\u281C':")",
            '\u2802':",",
            '\u2822':"?",
            '\u2800':" ",
    }

    var rc = map[rune]string{
            '\u2801':"А",
            '\u2803':"Б",
            '\u283A':"В",
            '\u281B':"Г",
            '\u2819':"Д",
            '\u2811':"Е",
            '\u2821':"Ё",
            '\u281A':"Ж",
            '\u2835':"З",
            '\u280A':"И",
            '\u282F':"Й",
            '\u2805':"К",
            '\u2807':"Л",
            '\u280D':"М",
            '\u281D':"Н",
            '\u2815':"О",
            '\u280F':"П",
            '\u2817':"Р",
            '\u280E':"С",
            '\u281E':"Т",
            '\u2825':"У",
            '\u280B':"Ф",
            '\u2813':"Х",
            '\u2809':"Ц",
            '\u281F':"Ч",
            '\u2831':"Ш",
            '\u282D':"Щ",
            '\u2837':"Ъ",
            '\u282E':"Ы",
            '\u283E':"Ь",
            '\u282A':"Э",
            '\u2833':"Ю",
            '\u282B':"Я",
            '\u2832':".",
            '\u2816':"!",
            '\u2824':"-",
            '\u2826':"\"",
            '\u2834':"\"",
            '\u2823':"(",
            '\u281C':")",
            '\u2802':",",
            '\u2822':"?",
            '\u2800':" ",

    }

    var s = map[string]rune{
            "А":'\u2801',
            "Б":'\u2803',
            "В":'\u283A',
            "Г":'\u281B',
            "Д":'\u2819',
            "Е":'\u2811',
            "Ё":'\u2821',
            "Ж":'\u281A',
            "З":'\u2835',
            "И":'\u280A',
            "Й":'\u282F',
            "К":'\u2805',
            "Л":'\u2807',
            "М":'\u280D',
            "Н":'\u281D',
            "О":'\u2815',
            "П":'\u280F',
            "Р":'\u2817',
            "С":'\u280E',
            "Т":'\u281E',
            "У":'\u2825',
            "Ф":'\u280B',
            "Х":'\u2813',
            "Ц":'\u2809',
            "Ч":'\u281F',
            "Ш":'\u2831',
            "Щ":'\u282D',
            "Ъ":'\u2837',
            "Ы":'\u282E',
            "Ь":'\u283E',
            "Э":'\u282A',
            "Ю":'\u2833',
            "Я":'\u282B',
            "A":'\u2801',
            "B":'\u2803',
            "C":'\u2809',
            "D":'\u2819',
            "E":'\u2811',
            "F":'\u280B',
            "G":'\u281B',
            "H":'\u2813',
            "I":'\u280A',
            "J":'\u281A',
            "K":'\u2805',
            "L":'\u2807',
            "M":'\u280D',
            "N":'\u281D',
            "O":'\u2815',
            "P":'\u280F',
            "Q":'\u281F',
            "R":'\u2817',
            "S":'\u280E',
            "T":'\u281E',
            "U":'\u2825',
            "V":'\u2827',
            "W":'\u283A',
            "X":'\u282D',
            "Y":'\u283D',
            "Z":'\u2835',
            ".":'\u2832',
            "!":'\u2816',
            "-":'\u2824',
            "\"":'\u2826',
            // "\"":'\u2834',
            "(":'\u2823',
            ")":'\u281C',
            ",":'\u2802',
            "?":'\u2822',
            " ":'\u2800',
    }

    return &bri{
        smap: s,
        rc: rc,
        rl: rl,
    }
}

func(b *bri) Encode(s string) string {
    var result string

    for _, char := range s {

        key := string(char)

        key = strings.ToUpper(key)

        v, ok := b.smap[key]

        if !ok {
            result += key
        }

        result += string(v)
    }

    return result

}

func(b *bri) Decode(s string) string {
    var result string
    var lib  map[rune]string

    if b.containsCirilics(s) {
        lib = b.rc
    } else {
        lib = b.rl
    }

    for _, char := range s{

        v, ok := lib[char]

        if !ok {
            result += string(char)
        }

        result += v
    }

    return result
}

func(b bri) containsCirilics(s string) bool {

    for key := range b.rl {
        delete(b.rc, key)
    }

    for _, char := range s{

        if _, ok := b.rc[char]; ok {
            return true
        }
    }

    return false
}

func(b bri) IsEncoded(s string) bool {
    var count int32

    for _, char := range s {
        count += char
    }


    count = count / int32(len(s))

    return count > 1000
}
