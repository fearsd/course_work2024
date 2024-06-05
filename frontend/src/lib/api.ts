import type { ResponseBody } from "./utils";
import axios from "axios";

const API_URL = 'http://localhost:3333';

export async function getRandomPermutation(arr: number[], edgeIndex: number, lastIndex: number): Promise<ResponseBody | null> {
    const data = { arr, edgeIndex, lastIndex };
    try {
        const response = await axios.post<ResponseBody>(`${API_URL}/algo`, data);
        return response.data;
    } catch (error) {
        console.log(error);
        return null;
    }
}