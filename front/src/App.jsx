import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { Home } from './views/Home'
import { About } from './views/About'
import { Auth } from './views/Auth'
import { NotFound } from './views/NotFound'
import { Dashboard } from './views/dashboard/Dashboard'
import { Settings } from './views/dashboard/Settings'
import DashboardLayout from './components/layouts/DashboardLayout'
import Pancake from './views/dashboard/Pancake'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/home" replace />} />
        <Route path="/home" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/auth" element={<Auth />} />
        <Route path="/dashboard" element={<DashboardLayout />}>
          <Route index element={<Dashboard />} />
          <Route path="settings" element={<Settings />} />
          <Route path="pancake" element={<Pancake />} />
        </Route>
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
