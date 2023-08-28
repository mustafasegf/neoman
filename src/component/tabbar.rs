use crate::items::Item;

#[derive(Debug, Default)]
pub struct TabBar {
    pub selected: usize,
    pub tabs: Vec<Item>,
}
