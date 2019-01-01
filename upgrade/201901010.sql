insert into `user_deck_rela` (uid, did, ready_count, new_count) select user_id, id, ready_count, new_count from deck
ALTER TABLE `deck` DROP `user_id`;
ALTER TABLE `deck` DROP `new_count`;
ALTER TABLE `deck` DROP `ready_count`;

update note, card set note.did=card.did where note.id=card.nid;