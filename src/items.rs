use std::{
    cell::RefCell,
    ops::{Deref, DerefMut},
    rc::Rc,
};
use tui_tree_widget::{TreeItem, TreeState};

#[derive(Debug, Default)]
pub struct ItemInner {
    pub name: String,
    pub selected: bool,
    pub active: bool,
}

impl ItemInner {
    pub fn new(name: &str) -> Self {
        ItemInner {
            name: name.to_string(),
            selected: false,
            active: false,
        }
    }
}

impl std::fmt::Display for ItemInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.name)
    }
}

#[derive(Debug, Clone, Default)]
pub struct Item(Rc<RefCell<ItemInner>>);

impl std::fmt::Display for Item {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let inner = self.0.borrow();
        write!(f, "{}", inner)
    }
}

impl Deref for Item {
    type Target = Rc<RefCell<ItemInner>>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Item {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

impl Item {
    pub fn new(name: &str) -> Self {
        Item(Rc::new(RefCell::new(ItemInner::new(name))))
    }
}

#[derive(Debug, Default)]
pub struct StatefulTree<'a> {
    pub state: TreeState,
    pub items: Vec<TreeItem<'a, Item>>,
}

impl<'a> StatefulTree<'a> {
    #[allow(dead_code)]
    pub fn new() -> Self {
        Self {
            state: TreeState::default(),
            items: Vec::new(),
        }
    }

    pub fn with_items(items: Vec<TreeItem<'a, Item>>) -> Self {
        Self {
            state: TreeState::default(),
            items,
        }
    }

    pub fn first(&mut self) {
        self.state.select_first();
    }

    pub fn last(&mut self) {
        self.state.select_last(&self.items);
    }

    pub fn down(&mut self) {
        self.state.key_down(&self.items);
    }

    pub fn up(&mut self) {
        self.state.key_up(&self.items);
    }

    pub fn left(&mut self) {
        self.state.key_left();
    }

    pub fn right(&mut self) {
        self.state.key_right();
    }

    pub fn toggle(&mut self) {
        self.state.toggle_selected();
    }

    pub fn selected(&self) -> Option<&TreeItem<'a, Item>> {
        let indicies = self.state.selected();
        let item = self.items.get(indicies[0]);

        indicies.iter().skip(1).fold(item, |item, &i| {
            item.and_then(|item| item.children().get(i))
        })
    }
}
