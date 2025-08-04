import { Outlet } from 'react-router-dom'
import { Navbar } from '../ui/Navbar'

const DashboardLayout = () => {
  return (
    <div className="dashboardContainer min-h-screen flex w-full h-full">
      <Navbar />
      <main className="flex w-full min-h-screen p-[10px] bg-gray-200">
        <div className="m-10px w-full h-full bg-white rounded-[20px] shadow-md">
          <Outlet />
        </div>
     </main>
    </div>
  )
}

export default DashboardLayout;

