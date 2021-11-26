const newArray = ["1", "2", "3", "3", 5, 6, 9, 4, true]
const exceptArray = ["2", "3", "2", 9, 11]

// 需求把newArray中的exceptArray去掉 -- 基础数据

function ArrayDrop(array: any[], ...drops: any[]) {
    return array.filter((it) => {
        return !drops.some((its) => {
            return its === it
        })
    })
}

const aa = ArrayDrop(newArray, ...exceptArray)
console.log(aa)

// 去重
let ax = new Set(newArray)
console.log(Array.from(ax))

