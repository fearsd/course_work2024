export function fromStringToArrayOfNum(stringToTransform: string): number[] {
    return stringToTransform.split(',')
            .map(s => s.trim())
            .filter(s => s !== '')
            .map(Number)
            .filter(n => !isNaN(n));
}

export type ResponseBody = {
    arr: number[],
    edgeIndex: number,
    lastIndex: number,
    randIndex: number
}
