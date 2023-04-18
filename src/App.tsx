import ReactDOM from 'react-dom/client'
import { useState, useEffect } from 'react'

const App = () => {
  const [message, setMessage] = useState('')

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetch('/api')
      console.log(res.status)
      console.log(res.body)
      const data = await res.json()
      console.log(data)
      setMessage(data.message)
      console.log("end fetch")
    }
    fetchData()
  }, [])

  return <h1>{message}</h1>
}

export default App