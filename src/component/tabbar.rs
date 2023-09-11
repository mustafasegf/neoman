use crate::items::Item;

#[derive(Debug, Default)]
pub struct TabBar {
    pub selected: usize,
    pub tabs: Vec<Item>,
}

impl TabBar {
    pub fn right(&mut self) {
        if self.selected < self.tabs.len() - 1 {
            self.selected += 1;
        }
    }
    pub fn left(&mut self) {
        if self.selected > 0 {
            self.selected -= 1;
        }
    }

    pub fn select(&mut self, index: usize) {
        self.selected = index;
    }

    pub fn first(&mut self) {
        self.selected = 0;
    }

    pub fn last(&mut self) {
        self.selected = self.tabs.len() - 1;
    }

    pub fn add(&mut self, item: Item) {
        self.tabs.push(item);
    }

}
