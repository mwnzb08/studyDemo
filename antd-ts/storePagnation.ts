// 前端模拟分页
const list = [] as string[]

for (let i = 0; i < 100000; i++) {
    list.push('Noo.' +i)
}

const Pages = {
    page_size:10,
    page_index: 0,
    page_total: 0,
}

const resultList = () => {
    let list2 =  list.slice(Pages.page_index, Pages.page_index + Pages.page_size)
    console.log(list2)
    setTimeout(() => {
        Pages.page_index = 10
        let list3 =  list.slice(Pages.page_index, Pages.page_index + Pages.page_size)
        list2.push(...list3)
        console.log(list2)
    },500)
}

resultList()


