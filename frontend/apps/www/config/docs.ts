import { MainNavItem, SidebarNavItem } from "types/nav"

interface DocsConfig {
    mainNav: MainNavItem[]
    sidebarNav: SidebarNavItem[]
}

export const docsConfig: DocsConfig = {
    mainNav: [
        {
            title: "Documentation",
            href: "/docs",
        },
        {
            title: "Components",
            href: "/docs/components/accordion",
        },
        {
            title: "Examples",
            href: "/examples",
        },
        {
            title: "Figma",
            href: "/docs/figma",
        },
        {
            title: "GitHub",
            href: "https://github.com/shadcn/ui",
            external: true,
        },
        {
            title: "Twitter",
            href: "https://twitter.com/shadcn",
            external: true,
        },
    ],
    sidebarNav: [
        {
            title: "Guides",
            items: [
                {
                    title: "Introduction",
                    href: "/docs",
                    items: [],
                },
                {
                    title: "Quickstart",
                    href: "/docs/installation",
                    items: [],
                },
                {
                    title: "SDKs",
                    href: "/docs/theming",
                    items: [],
                },
                {
                    title: "Authentication",
                    href: "/docs/cli",
                    items: [],
                },
                {
                    title: "Pagination",
                    href: "/docs/components/typography",
                    items: [],
                },
                {
                    title: "Errors",
                    href: "/docs/components/typography",
                    items: [],
                },
                {
                    title: "Webhooks",
                    href: "/docs/components/typography",
                    items: [],
                },
            ],
        },
        {
            title: "Run Nodes",
            items: [
                {
                    title: "Token Economics",
                    href: "/docs/token",
                    items: [],
                },
                {
                    title: "Join Testnet",
                    href: "/docs/join-network",
                    items: [],
                },
            ],
        },
        {
            title: "Build Apps",
            items: [
                {
                    title: "Identities",
                    href: "/docs/reference/identity",
                    items: [],
                },
                {
                    title: "Services",
                    href: "/docs/reference/services",
                    items: [],
                },
                {
                    title: "Wallets & Accounts",
                    href: "/docs/reference/services",
                    items: [],
                },
            ],
        },
    ],
}
