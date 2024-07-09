import Header from "../components/Header"
import HeaderMobile from "../components/HeaderMobile"
import MarginWidthWrapper from "../components/MarginWidthWrapper"
import PageWrapper from "../components/PageWrapper"
import SideNav from "../components/SideNav"

export default function ManagementLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <section>
      <div className="flex">
        <SideNav />
        <main className="flex-1">
          <MarginWidthWrapper>
            <Header />
            <HeaderMobile />
            <PageWrapper>{children}</PageWrapper>
          </MarginWidthWrapper>
        </main>
      </div>
    </section>
  )
}
