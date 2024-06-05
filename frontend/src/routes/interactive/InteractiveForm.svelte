<script lang="ts">
	import { getRandomPermutation } from '$lib/api';
    import { fromStringToArrayOfNum } from '$lib/utils';
	import Iteration from './Iteration.svelte';

    type Iter = {
        leftSide: string,
        rightSide: string,
        edgeNum: number,
        randNum: number
    }

    let seq: string = "1, 2, 3, 4";
    let ans: string = "";
    let err: string = "";
    let isStarted = false;
    let isEnd = false;
    let arr: number[];
    let lastIndex: number;
    let edgeIndex: number;
    let iterations: Iter[] = [];

    function addIteration(arr: number[], edgeIndex: number, randIndex: number): void {
        iterations.push({leftSide: arr.slice(0, edgeIndex).join(", "), rightSide: arr.slice(edgeIndex).join(", "), edgeNum: arr[edgeIndex], randNum: arr[randIndex]})
        iterations = iterations;
    }

    async function generatePermutation(event: Event): Promise<void> {
        arr = fromStringToArrayOfNum(seq);
        edgeIndex = arr.length - 1;
        lastIndex = arr.length - 2;
        if (arr.length === 0) {
            err = "Введена пустой список";
            ans = "";
            return;
        }

        try {
            const response = await getRandomPermutation(arr, edgeIndex, lastIndex);
            if (response && response.arr) {
                arr = response.arr;
                addIteration(arr, edgeIndex, response.randIndex);
                isStarted = !isStarted;
            } else {
                throw new Error("Некорректный ответ от сервера");
            }
        } catch (error) {
            console.error(error);
            ans = "";
            err = "Произошла ошибка: " + (error instanceof Error ? error.message : error);
        }
    }

    async function triggerNextStep(event: Event): Promise<void> {
        edgeIndex -= 1;
        lastIndex -= 1;
        if (lastIndex==-1 && edgeIndex==0) {
            isEnd = true;
            
            ans = arr.slice(edgeIndex).join(", ");
            return
        }
        try {
            // console.log(arr, edgeIndex, lastIndex);
            const response = await getRandomPermutation(arr, edgeIndex, lastIndex);
            if (response && response.arr) {
                arr = response.arr;
                addIteration(arr, edgeIndex, response.randIndex);
                console.log(response);
            } else {
                throw new Error("Некорректный ответ от сервера");
            }
        } catch (error) {
            console.error(error);
            ans = "";
            err = "Произошла ошибка: " + (error instanceof Error ? error.message : error);
        }
    }
   
</script>

<form action="">
    <label for="">Введите список чисел:</label>
    <input type="text" bind:value={seq}>
    <button on:click={generatePermutation} disabled={isStarted}>Сгенерировать случайную перестановку</button>
</form>

{#if isStarted}
<button on:click={triggerNextStep} disabled={isEnd}>Следующий шаг</button>

{#each iterations as iter, index}
    <Iteration {...iter}/>
{/each}
{/if}

{#if isEnd}
<h3>Итоговая перестановка:</h3>
<p>{ans}</p>
{/if}