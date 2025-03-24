import {NavLink} from "react-router-dom";

const NavLinks = () => {
    return <>
        <NavLink to="/">Home</NavLink>
        <NavLink to="/blog">Blog</NavLink>
    </>
}

const Navbar = () => {
    return <>
        <nav className="w-1/8 flex justify-end">
            <div className="flex w-full justify-between">
                <NavLinks />
            </div>
        </nav>
    
    
    </>
}

export default Navbar
