import axios from "../../utils/singletons/axios";

export const getBudgets = async ({id}: {id: string}) => {
    const response = await axios.get("/budgets/" + id);
    return response.data;    
}

