use crate::items::StatefulTree;

#[derive(Debug)]
pub struct SideBar {
    pub size: u16,
    pub selected: usize,
    pub tree: StatefulTree<'static>,
}
