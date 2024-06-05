<script lang="ts">
	import { getRandomPermutation } from '$lib/api';
    import { fromStringToArrayOfNum } from '$lib/utils';
    let seq: string = "1, 2, 3, 4";
    let ans: string = "";
    let err: string = "";

    async function generatePermutation(event: Event): Promise<void> {
        let arr: number[] = fromStringToArrayOfNum(seq);
        if (arr.length === 0) {
            err = "Введена пустой список";
            ans = "";
            return;
        }

        try {
            const response = await getRandomPermutation(arr, arr.length - 1, 0);
            if (response && response.arr) {
                ans = response.arr.join(", ");
                err = "";
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
    <button on:click={generatePermutation}>Сгенерировать случайную перестановку</button>
</form>

{#if ans != ""}
    <p>Результат: {ans}</p>
{/if}

{#if err != ""}
    <p>{err}</p>
{/if}
