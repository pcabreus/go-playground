let count = 0

const button = document.getElementById('button') as HTMLButtonElement
const counterValue = document.querySelector('.counter-value') as HTMLSpanElement
counterValue.textContent = count.toString()

button.onclick = () => {
    count++
    counterValue.textContent = count.toString()    
}

let container: string[] = []
function solution(queries: string[][]): string[] {
    let result: string[] = []
    for (let i = 0; i < queries.length; i++) {
        let action = queries[i][0]
        let number = queries[i][1]
        
        switch (action) {
            case 'add':
                container[parseInt(number)] = number
                break
            case 'get':
                let count = container.filter(x => x === number).length
                result.push(count.toString())
                break
        }
    } 
}