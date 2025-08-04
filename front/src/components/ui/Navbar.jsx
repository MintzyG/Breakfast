import { useState } from 'react'
import { NavLink } from 'react-router-dom'
import { Menu, X, Home, Settings } from 'lucide-react'

export const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <aside className={`bg-white shadow-md ${isOpen ? 'w-64' : 'w-16'} transition-all duration-300`}>
      <button
        className="p-4 focus:outline-none"
        onClick={() => setIsOpen(!isOpen)}
        aria-label="Toggle Menu"
      >
        {isOpen ? <X /> : <Menu />}
      </button>

      <nav className="mt-4 space-y-2">
        <NavLink
          to="/dashboard"
          end
          className={({ isActive }) =>
            `flex items-center px-4 py-2 rounded-md ${isActive ? 'bg-blue-50 text-blue-600' : 'text-gray-600 hover:bg-gray-50'}`
          }
        >
          <span className="mr-2">
            {/* Icon for Overview */}
            <Home />
          </span>
          {isOpen && <span>Overview</span>}
        </NavLink>
        <NavLink
          to="/dashboard/settings"
          className={({ isActive }) =>
            `flex items-center px-4 py-2 rounded-md ${isActive ? 'bg-blue-50 text-blue-600' : 'text-gray-600 hover:bg-gray-50'}`
          }
        >
          <span className="mr-2">
            {/* Icon for Settings */}
            <Settings />
          </span>
          {isOpen && <span>Settings</span>}
        </NavLink>
        <NavLink
          to="/dashboard/pancake"
          className={({ isActive }) =>
            `flex items-center px-4 py-2 rounded-md ${isActive ? 'bg-blue-50 text-blue-600' : 'text-gray-600 hover:bg-gray-50'}`
          }
        >
          <span className="mr-2">
          {/* Icon for Settings */}
          <Home />
          </span>
          {isOpen && <span>Settings</span>}
        </NavLink>
      </nav>
    </aside>
  )
}
