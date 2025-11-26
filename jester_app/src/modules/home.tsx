import { useEffect, useState } from "react";

//TODO: refactor into separate files
interface budgetType {
    id: Number,
    name: String,
    allocated: Number,
    spent: Number,
    color: String,
} 

interface entryType {
    id: Number | null,
    description: String,
    date: String,
    amount: Number,
    budgetId: Number,
}

const seed_budgets: budgetType[] = [
    { id: 1, name: "Groceries", allocated: 86.60, spent: 86.60, color: "#d9d2e9"},
    { id: 2, name: "Disgressionary Food", allocated: 45.00, spent: 34.50, color: "#d9d2e9"},
    { id: 3, name: "Gym", allocated: 45.00, spent: 0, color: "#d9d2e9"},
    { id: 4, name: "Car Gas", allocated: 100.00, spent: 100, color: "#d9d2e9"}
]



const Home = () => {

     // const instance = axios.create({
  //   baseURL: "http://localhost:8080",
  //   timeout: 1000,
  //   withCredentials: false,
  //   headers: {
  //     "Content-Type" : "application/json",
  //   }
  // })

  const [budgets, setBudgets] = useState<budgetType[]>(seed_budgets);
  const [entries, setEntries] = useState<entryType[]>([]);

  let totalAllocated: Number = 0;
  let totalSpent: Number = 0;
  let totalRemaining: Number = 0;

  for(let i = 0; i < budgets.length; i++) {
    totalAllocated = Number(totalAllocated) + Number(budgets[i].allocated);
    totalSpent = Number(totalSpent) + Number(budgets[i].spent);
  }

  totalRemaining = Number(totalAllocated) - Number(totalSpent);


  const handleNewEntry = () => {
    return (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      const formData = new FormData(event.currentTarget);
      const description = formData.get("description") as string;
      const date = formData.get("date") as string;
      const amount = formData.get("amount") as string;
      const budget = formData.get("budget") as string;

      //inserting the new entry logic here
      const newEntry: entryType = {
        id: null,
        description: description,
        date: date,
        amount: Number(amount),
        budgetId: Number(budget),
      }

      setEntries([...entries, newEntry]);

      console.log(description, date, Number(amount).toFixed(2), budget);
    }
  }

  useEffect(() => {
    
  },[budgets, entries])

    return (
        <div style={{
      display: "flex",
      flexDirection: "column",
    }}>

      <h1>Budget Overview</h1>
      <table>
        <thead>
          <tr>
            <th style={{textAlign: "left"}}>Budget</th>
            <th style={{textAlign: "left"}}>Allocated</th>
            <th style={{textAlign: "left"}}>Spent</th>
            <th style={{textAlign: "left"}}>Remaining</th>
          </tr>
        </thead>
        <tbody>
          {budgets.map((budget) => {
            return (
              <tr key={budget.id.toString()}>
                <td>{budget.name}</td>
                <td>${budget.allocated.toFixed(2)}</td>
                <td>${budget.spent.toFixed(2)}</td>
                <td>${(Number(budget.allocated) - Number(budget.spent)).toFixed(2)}</td>
              </tr>
            )
          })}
        </tbody>
        <tfoot>
          <tr>
            <td>Total</td>
            <td>${totalAllocated.toFixed(2)}</td>
            <td>${totalSpent.toFixed(2)}</td>
            <td>${totalRemaining.toFixed(2)}</td>
          </tr>
        </tfoot>
      </table>

      <div>
        <h2>Budget Re-allocation</h2>
        <label>Amount:</label>
        <input type="number" name="reallocation-amount" step="0.01" />
        <br/>
        <label>From Budget:</label>
        <select name="from-budget">
          {budgets.map((budget) => {
            return (
              <option key={budget.id.toString()} value={budget.id.toString()}>{budget.name}</option>
            )
          })}
        </select>
        <p>put new budget amount here</p>
        <br/>
        <label>To Budget:</label>
        <select name="to-budget">
          {budgets.map((budget) => {
            return (
              <option key={budget.id.toString()} value={budget.id.toString()}>{budget.name}</option>
            )
          })}
        </select>
        <p>put new budget amount here</p>
        <br/>
        <button type="button">Re-allocate</button>
      </div>

      <br/>

      <h1>Budget Entries</h1>
      <table>
        <thead>
          <th>Description</th>
          <th>Date</th>
          <th>Amount</th>
          <th>Budget</th>
        </thead>
        <tbody>
          {entries.map((entry) => {
            return (
              <tr key={entry.id?.toString()}>
                <td>{entry.description}</td>
                <td>{entry.date}</td>
                <td>${entry.amount.toFixed(2)}</td>
                <td>{budgets.find(budget => budget.id === entry.budgetId)?.name}</td>
              </tr>
            )
          })}
        </tbody>
        <tfoot></tfoot>
      </table>

      <form onSubmit={handleNewEntry()}>
        <h2>Add Budget Entry</h2>
        <label>Description:</label>
        <input type="text" name="description" />
        <br/>
        <label>Date:</label>
        <input type="date" name="date" />
        <br/>
        <label>Amount:</label>
        <input type="number" name="amount" step="0.01" />
        <br/>
        <label>Budget:</label>
        <select name="budget">
          {budgets.map((budget) => {
            return (
              <option key={budget.id.toString()} value={budget.id.toString()}>{budget.name}</option>
            )
          })}
        </select>
        <br/>
        <button type="submit">Add Entry</button>
      </form>

    </div>
    )
}

export default Home;