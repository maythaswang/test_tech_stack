import Logo from "./Logo";
import Navbar from "./Navbar";

function Header() {
  return (
    <header className="sticky top-0 z-[1] mx-auto flex w-full max-w-8xl flex-wrap items-center justify-between border-b border-gray-100 bg-background p-[2em] font-sans font-bold uppercase text-text-primary backdrop-blur-[100px] dark:border-gray-800 dark:bg-d-background dark:text-d-text-primary">
      <Logo />
      <h3 className="absolute left-1/2 transform -translate-x-1/2 text-center">REACT + POSTGRES + GO</h3>
      <Navbar />
    </header>
  );
}
export default Header;
