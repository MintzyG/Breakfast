import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '../../../axios'

const AuthTabs = () => {
  const [activeTab, setActiveTab] = useState('login')
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
  })
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    })
  }

  const handleSubmit = async (e, type) => {
    e.preventDefault()
    setError('')
    try {
      localStorage.removeItem('jwt')

      const endpoint = type === 'login' ? '/login' : '/register'
      const response = await api.post(endpoint, formData)

      console.log(`${type} response:`, response.data)

      const token = response.data

      if (!token) {
        console.error('No token received from server')
        setError('Authentication error: No token received')
        return
      }

      console.log(`Token received:`, token.substring(0, 20) + '...')

      try {
        await api.post('/verify-jwt', {}, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })

        localStorage.setItem('jwt', token)
        navigate('/dashboard')
      } catch (verifyErr) {
        console.error('Token verification failed:', verifyErr.response?.data)
        setError(verifyErr.response?.data?.message || 'Token verification failed')
      }
    } catch (err) {
      console.error(`${type} request failed:`, err.response?.data)
      setError(err.response?.data?.message || 'Something went wrong')
    }
  }

  return (
    <div className="w-full">
      <div className="flex border-b mb-4">
        <button
          className={`px-4 py-2 ${
            activeTab === 'login'
              ? 'border-b-2 border-blue-500 text-blue-600'
              : 'text-gray-600'
          }`}
          onClick={() => setActiveTab('login')}
        >
          Login
        </button>
        <button
          className={`px-4 py-2 ${
            activeTab === 'signup'
              ? 'border-b-2 border-blue-500 text-blue-600'
              : 'text-gray-600'
          }`}
          onClick={() => setActiveTab('signup')}
        >
          Sign Up
        </button>
      </div>

      {error && <p className="text-red-500">{error}</p>}

      {activeTab === 'login' ? (
        <form className="space-y-4" key="login" onSubmit={(e) => handleSubmit(e, 'login')}>
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <input
              type="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Password</label>
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-600 text-white rounded-md py-2 hover:bg-blue-700"
          >
            Login
          </button>
        </form>
      ) : (
        <form className="space-y-4" key="signup" onSubmit={(e) => handleSubmit(e, 'register')}>
          <div>
            <label className="block text-sm font-medium text-gray-700">Name</label>
            <input
              type="text"
              name="name"
              value={formData.name}
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <input
              type="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Password</label>
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleChange}
              className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-600 text-white rounded-md py-2 hover:bg-blue-700"
          >
            Sign Up
          </button>
        </form>
      )}
    </div>
  )
}

export default AuthTabs

