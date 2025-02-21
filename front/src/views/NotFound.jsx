import { NavLink } from 'react-router-dom'

export const NotFound = () => (
  <div className="min-h-screen bg-gray-50 flex items-center justify-center">
    <div className="bg-white rounded-lg shadow p-8 text-center">
      <h1 className="text-6xl font-bold text-gray-900">404</h1>
      <p className="mt-4 text-gray-600">Page not found</p>
      <NavLink to="/" className="mt-6 inline-block text-blue-600 hover:text-blue-700">
        Go back home
      </NavLink>
    </div>
  </div>
)
