
import { ModeToggle } from "@/components/mode-toggle";
import { NavigationMenuDemo } from "@/components/navigation-menu";
import { SignupForm } from "@/components/signup";
import { LoginForm } from "@/components/login";
export default function Home() {
  return (
    <main>
      <NavigationMenuDemo/>
      {/* <ModeToggle /> */}
      <SignupForm />
      <LoginForm/>
      This is main
    </main>
  );
}
