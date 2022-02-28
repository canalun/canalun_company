declare module "*.jpg";
declare module "*.png";
declare module "*.jpeg";
declare module "*.gif";

declare module "@materials/entry_list/*.json" {
    interface Entry {
        title: string
        url: string
        lastUpdatedAt: string
    }

    const data: Entry[]

    export default data
}