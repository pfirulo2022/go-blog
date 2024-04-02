import { Routes, Route } from "react-router-dom";
import Home from './page/Home';
import Blog from './page/Blog';
import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";
import Add from "./page/Add";
import Edit from "./page/Edit";
import Delete from "./page/Delete";
import Contact from "./page/Contact";

function App() {

  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/add" element={<Add />} />
        <Route path="/edit/:id" element={<Edit />} />
        <Route path="/delete/:id" element={<Delete />} />
        <Route path="/blog/:id" element={<Blog />} />
        <Route path="/contact" element={<Contact />} />
        {/*   <Route path="/login" element={<Login />} />
        <Route path="/logout" element={<Logout />} />
        <Route path="/about" element={<About />} />
         */}
      </Routes>
      <Footer />
    </>
  );
}

export default App
