use tui_tree_widget::TreeItem;

use crate::items::{Item, StatefulTree};

#[derive(Debug)]
pub struct SideBar {
    pub size: u16,
    pub selected: usize,
    pub tree: StatefulTree<'static>,
}

impl SideBar {
    pub fn selected(&self) -> Option<&TreeItem<Item>> {
        let indicies = self.tree.state.selected();
        let item = self.tree.items.get(indicies[0]);

        indicies.iter().skip(1).fold(item, |item, &i| {
            item.and_then(|item| item.children().get(i))
        })
    }
}
