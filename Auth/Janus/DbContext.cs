using Janus.Identity;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;

namespace Janus
{
    public class DbContext : IdentityDbContext<User>
    {
        public DbContext(DbContextOptions options) : base(options)
        {
        }

        protected override void OnModelCreating(ModelBuilder builder)
        {
            base.OnModelCreating(builder);
            builder.ApplyConfiguration(new RoleConfiguration());
        }

        

        public DbSet<RefreshToken> RefreshToken => Set<RefreshToken>();
        public DbSet<Profile> Profiles { get; set; }

    }
}
