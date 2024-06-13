import "./globals.css";

import { ThemeProvider } from "@/components/theme-provider";
import { NavigationMenuDemo } from "@/components/navigation-menu";
export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode; }>) {
  return (
    <>
      <html lang="en" suppressHydrationWarning>
        <head />
        <body>
          <ThemeProvider
            attribute="class"
            defaultTheme="system"
            enableSystem
            disableTransitionOnChange
          >
            <NavigationMenuDemo />
            {children}
          </ThemeProvider>
        </body>
      </html>
    </>
  )
}
