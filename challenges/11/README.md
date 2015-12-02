## Challenge for 12/11: Rebasing

Today we will examine the [git-rebase](https://git-scm.com/docs/git-rebase) command, commonly invoked as `git rebase`. What does this command do?
````
Forward-port local commits to the updated upstream head
````
So, the documentation isn't particularly helpful in describing this function to a neophyte. We can better describe a rebase as allowing us to rewrite our commit history. There are many things we can do with the rebase command. We will focus on two particular use cases today.

### Re-apply your branches commits on an updated master

The documentation text is not awesome but there are some useful diagrams. Imagine that we created a branch `topic` some time ago from `master`, but additional commits have been made since then.
````
          A---B---C topic
         /
    D---E---F---G master
````
The commits `F` or `G` (equivalent to `HEAD`) could have modified the same file segments as `A-C`. Even without conflicts, our commits could be incompatible or contradictory to the other changes. We want to reapply our series of commits against `HEAD` rather than `E`. If we use the command `git rebase -i origin/master`, our history will now look like this:
````
                  A'--B'--C' topic
                 /
    D---E---F---G master
````
Voila! Our commits are applied against `HEAD`. We can now run our unit tests (we do have unit tests, *right*?) and ensure everything works properly before continuing our work.

If we run into conflicts, where commits `F-G` conflicted with the commits in `A-C`, we will be presented with the opportunity to resolve the conflict before proceeding. After editing the file in question to resolve the conflict, we `git add` the file and `git rebase --continue` to proceed. We can back out of the rebase with `git rebase --abort`.

One final note: When we run rebase, we use the format *[remote/]branch*, where the remote is optional. I encourage you to use the remote when working on GitHub projects, as your local branch is potentially older than the remote's. Always be sure to fetch from the remote (`git fetch *remote*`) as well.

### Rewriting history

Another cool trick is to use rebase to rewrite our history. Technically, rebasing against master is rewriting the history as well, but I'm talking more specifically about modifying a commit itself. We perform the rebase the same way (`git rebase -i origin/master`) but we're not interested in applying things as-is so we'll look at the options available to us:
````
# Commands:
#  p, pick = use commit
#  r, reword = use commit, but edit the commit message
#  e, edit = use commit, but stop for amending
#  s, squash = use commit, but meld into previous commit
#  f, fixup = like "squash", but discard this commit's log message
#
# If you remove a line here THAT COMMIT WILL BE LOST.
# However, if you remove everything, the rebase will be aborted.
````
Each commit defaults to `pick`. Let's break up the available options into four  groups:

#### Preserve commits

If we select `pick` or `reword`, the commit itself is left alone. When we use `reword`, we are given the opportunity to edit the message attached to the commit. These are the least intrusive rewrite options.

#### Edit commits

If we select `edit`, we are given the ability to modify that commit before proceeding. In our example of `A-B-C` commits, if we edit `B`, we will essentially `git checkout` B's reference, then have the ability to make further changes. During this interval, `C` is not yet applied. After we continue the rebase, `C` would then be applied. This is mildly intrusive and we need to be careful not to create a conflict with future commits.

#### Merge commits

The `squash` and `fixup` options really do rewrite the history. If we `squash` commit `B`, we would be left with:
````
                  (A+B)'--C' topic
                 /
    D---E---F---G master
````
Commits `A` and `B` are now represented by a single commit that consists of all the changes in the previous two commits. The two options let us choose whether we want to keep the commit message (which we will get to edit as well) or discard it. I find `fixup` is very helpful when the squashed commits have messages like *typo* and *more typos*, and `squash` is more appropriate when the commit represents a significant but non-atomic change.

These are considered some of the most intrusive changes. Some people prefer to keep a 100% accurate history, though any given commit reference may be in an unusable state. Others prefer that each commit represent a single, atomic delta. Many fall in the middle of these ranges. Ensure that you discuss the rebase preferences with other project team members and that you follow the guidance of the project owner when contributing to other people's projects.

#### Discard or re-order commits

The rebase dialog shows us one more option:
````
# If you remove a line here THAT COMMIT WILL BE LOST.
````
There are actually **two** options here. You can delete a line and that commit will be remove entirely. This is helpful when you've made some change that isn't necessary at all, especially if it's the last commit. Just delete the line and continue the rebase and poof, it's gone.

The other option that isn't shown is re-ordering commits. Simply cut a whole line and insert it in a different place and the commit will be applied at that point in the chain. You must again ensure that you do not create a conflict.
````
$ git log --pretty=oneline | head -2
f81c50c8e5c5dc0b8980a17fc7fd71a68bd699ff Removed line B
65798764a399357e2b0933350ec1270f58b8ef1d Removed line A
$ git rebase -i origin/master
# Original content:
pick e7059f3 Removed line A
pick 648d162 Removed line B

# New content:
pick 648d162 Removed line B
pick e7059f3 Removed line A

Successfully rebased and updated refs/heads/bogus.
$ git log --pretty=oneline | head -2
e7059f334c3b0b2120deb2c1a73bdfd278389579 Removed line A
648d162e0a64fa359ab2d741dd0fd21fcde7fe0e Removed line B
````

This is very much the most intrusive rebase capability. Always be careful when removing or re-ordering the commits, lest we need to turn to [git-reflog](https://git-scm.com/docs/git-reflog) for help with recovery!

### Sharing the new history

Don't forget to push the modified branch to the remote! You may have noted that our three commits above,`A-C`, became `A'-C'`, and that the checksums during re-ordering changed. Each commit is rewritten entirely, which generates a new checksum. When we want to push our branch to a remote that had `A-C` commits already, we will need to force it with `git push origin topic -f`.

Be very careful with force, it can be destructive! I prefer to push without the `-f` flag and receive an error indicating the checksums have changed, then hit the up arrow and add `-f` to the end of the command, to ensure I don't make a mistake like leaving out the branch name and using `git push -f`. If you create such a git-astrophe, you'll want to look into [git-reflog](https://git-scm.com/docs/git-reflog) for assistance in recovering from your error.
