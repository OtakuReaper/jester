import axios from "../../utils/singletons/axios";

export const getBudgets = async ({id}: {id: string}) => {
    const response = await axios.get("/budgets/" + id);
    return response.data;    
}

export const getEntries = async ({id}: {id: string}) => {
    const response = await axios.get("/entries/" + id);
    return response.data;
}

export const getPeriods = async ({id}: {id: string}) => {
    const response = await axios.get("/periods/" + id);
    return response.data;
}