<script lang="ts">
    export let leftSide;
    export let rightSide;
    export let edgeNum;
    export let randNum;
</script>

{#if leftSide == ''}
<p>{rightSide}</p>
{:else}
<p>{leftSide} | {rightSide}</p>
{/if}
{#if edgeNum !== randNum}
<p>Элементы {edgeNum} и {randNum} поменялись местами</p>
{:else}
<p>Рандомный индекс совпал с краевым, замены мест нет</p>
{/if}
<hr/>